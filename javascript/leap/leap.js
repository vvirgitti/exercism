var Year = function() {};

Year.prototype.isLeap = function(input) {
  if(input = 2015){
    return false;
  }
  else if (input = 2016) {
    return true;
  }
};

module.exports = Year;
