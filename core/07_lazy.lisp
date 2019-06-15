;;;; ale core: lazy sequences

(defmacro lazy-seq body
  `(lazy-seq* (lambda [] ,@body)))

(define (take count coll)
  ((fn take-inner [count coll]
     (lazy-seq
       (if (and (> count 0) (!empty? coll))
           (cons (first coll) (take-inner (dec count) (rest coll)))
           '())))
    count coll))

(define (take-while pred coll)
  (lazy-seq
    (when-let [s (seq coll)]
      (let [fs (first s)]
        (when (pred fs)
              (cons fs (take-while pred (rest s))))))))

(define (drop count coll)
  (lazy-seq
    ((fn drop-inner [count coll]
       (if (> count 0)
           (drop-inner (dec count) (rest coll))
           coll))
      count coll)))

(defn partition
  ([count coll]
    (partition count count coll))

  ([count step coll]
    (lazy-seq
      (when (seq? coll)
            (cons (to-list (take count coll))
                  (partition count step (drop step coll)))))))

(defn range
  ([]
    (range 0 '() 1))

  ([last]
    (range 0 last (if (> last 0) 1 -1)))

  ([first last]
    (if (> last first)
        (range first last 1)
        (range last first -1)))

  ([first last step]
    (let [cmp (cond (null? last) (constantly #t)
                    (< step 0)  >
                    :else       <)]
      (if (cmp first last)
          (cons first (lazy-seq (range (+ first step) last step)))
          []))))

(defn map
  ([func coll]
    ((fn map-single [coll]
       (lazy-seq
         (when (seq coll)
               (cons (func (first coll))
                     (map-single (rest coll))))))
      coll))

  ([func coll . colls]
    ((fn map-parallel [colls]
       (lazy-seq
         (when (apply true? (map !empty? colls))
               (let [f (to-vector (map first colls))
                     r (map rest colls)]
                 (cons (apply func f) (map-parallel r))))))
      (cons coll colls))))

(define (filter func coll)
   (lazy-seq
     ((fn filter-inner [coll]
        (when (seq coll)
              (let [f (first coll)
                    r (rest coll)]
                (if (func f)
                    (cons f (filter func r))
                    (filter-inner r)))))
      coll)))

(define (cartesian-product . colls)
  (let* [rotate-row
         (fn rotate-row [row orig-row]
           (if (seq row)
               (let [res (rest row)]
                 (if (seq res)
                     [#f res]
                     [#t orig-row]))
               [#t orig-row]))

         rotate-rest
         (fn rotate-rest [rest orig]
           (let [f  (first rest)
                 fo (first orig)
                 r  (rest rest)
                 ro (rest orig)]
             (if (seq r)
                 (let [res (rotate-rest r ro)]
                   (if (res 0)
                       (let [rr (rotate-row f fo)]
                         [(rr 0) (cons (rr 1) (res 1))])
                       [#f (cons f (res 1))]))
                 (let [res (rotate-row f fo)]
                   [(res 0) (list (res 1))]))))

         rotate
         (fn rotate [work orig]
           (let [res (rotate-rest work orig)]
             (unless (res 0) (res 1))))

         iter
         (fn iter [work]
           (let [f (to-vector (map first work))
                 r (rotate work colls)]
             (if r
                 (cons f (lazy-seq (iter r)))
                 (list f))))]
    (lazy-seq
      (when (apply false? (map empty? colls))
            (iter colls)))))

(defmacro for
  [seq-exprs . body]
  (assert-args
    (vector? seq-exprs)        "for-each bindings must be a vector"
    (even? (length seq-exprs)) "for-each bindings must be paired")
  (let* [args (gensym "args")

         split-bindings
         (fn split-bindings
           ([idx name coll]
            [(list name (list args idx))
             (list coll)])
           ([idx name coll . rest]
            (let [res (apply split-bindings (cons (inc idx) rest))]
              [(cons* (res 0) (list args idx) name)
               (cons coll (res 1))])))

         split (apply split-bindings (cons 0 seq-exprs))
         bind# (to-vector (split 0))
         seqs# (split 1)]
    `(map
       (lambda [,args] (let ,bind# ,@body))
       (cartesian-product ,@seqs#))))

(defmacro for-each
  [seq-exprs . body]
  `(last! (for ,seq-exprs ,@body)))

(define (concat . colls)
  ((fn concat-inner [colls]
     (lazy-seq
       (when (seq colls)
             (let [f (first colls)
                   r (rest colls)]
               (if (seq f)
                   (cons (first f)
                         (concat-inner (cons (rest f) r)))
                   (concat-inner r))))))
     colls))
