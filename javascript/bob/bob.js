
var Bob = function() {};

Bob.prototype.hey = function(sentence) {
    if (sentence.match(/[A-Z]{4,}/g)) {
        return 'Whoa, chill out!'
    } else if (sentence.match(/[A-Z]{2,}[!]$/g)) {
        return 'Whoa, chill out!'
    } else if (sentence.match(/[?]$/g)) {
        return 'Sure.'
    } else if (sentence.match(/^(?![\s\S])/g)) {
        return 'Fine. Be that way!'
    } else if (sentence.match(/^\s/g)) {
        return 'Fine. Be that way!'
    } else {
        return 'Whatever.'
    }

};

module.exports = Bob;
