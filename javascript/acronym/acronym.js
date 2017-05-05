function parse(phrase) {
  const wordList = phrase.match(/(\w+|[A-Z][a-z]+)/g);
  return wordList.map(w => {
    return w[0].toUpperCase()
  }).join('')

}

module.exports = {parse: parse};
