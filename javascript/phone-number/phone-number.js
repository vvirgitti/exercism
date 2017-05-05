function PhoneNumber(phone) {
  this.number = function () {
    const trimmed = phone.replace(/\s+/g, '');
    const sanitized =  trimmed.trim().replace(/[()-.]+/g, '');
    return verify(sanitized);
  };

  this.areaCode = function(){
    const numberInArray = phone.split('');
    return selectNumber(numberInArray, 0, 3);
  };

  this.toString = function(){
    const numberInArray = phone.split('');
    const area = selectNumber(numberInArray, 0, 3);
    const second = selectNumber(numberInArray, 3, 6);
    const end = selectNumber(numberInArray, 6, 10);

    return `(${area}) ${second}-${end}`;

  };
}

function verify(number) {
  const numberInArray = number.split('');

  if (numberInArray.length === 10) {
    return number;
  } else if (numberInArray.length === 11) {
    if (numberInArray[0] != '1') {
          return '0000000000'
        } else {
          numberInArray.shift();
          return numberInArray.join('');
        }
  } else {
    return '0000000000';
  }

}

function selectNumber(array, first, last) {
  return array.slice(first, last).join('');
}

module.exports = PhoneNumber;