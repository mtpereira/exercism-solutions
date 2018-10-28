class Pangram {
    companion object {
        private var counter: HashMap<Char, Int> = HashMap()

        fun isPangram(sentence: String): Boolean {
            cleanInput(sentence).forEach {
                counter[it.toLowerCase()] = 1
            }
            return 26 == counter.values.reduce { acc, i -> acc + i }
        }

        private fun cleanInput(sentence: String): String {
            val ret = sentence.replace("[^\\w]".toRegex(), "")
            println(ret)
            return ret
        }
    }



}
