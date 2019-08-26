# printing
p "hello world"
p "%d %s" % [3, "rubies"]

# simple iteration
3.times { p "Hello!" }
0.upto(2) {|x| p x }

# arrays
ages = [5, 93, 12]
ages.each do |x|
  p "testing #{x}"
end

# map! applies changes in place
ages.map!{|x| x + 10}
adults = ages.select{|age| age >= 18 }
p "adults: #{adults}"

# hashes 
h = {
  :alpha => 1,
  :beta => 2,
}

h.each do |k,v|
  p "#{k}: #{v}"
end

# operators
x = 3
y = 4
p "min between #{x} and #{y} is %d" % [x < y ? x : y]


# methods
def yell(phrase)
  p phrase.upcase + "!"
end

yell "man, that ginger ale really hit the spot"

## the last expression is the return value
def sum(x, y)
  x + y
end

def isDoubleGreater(x, y)
  sum = x * 2
  [sum, sum > y]
end

p sum(1, 2)
res, isGreater = isDoubleGreater(2, 3)
p "#{res} #{isGreater}"

# gobal vars
$size = 5
p $size

# range and switch
foo = case $size
  when 0..2 then "small"
  when 3..5 then "medium"
  when 5..7 then "large"
  else nil
  end

p foo

def are_you_sure?
  while true
    print "Are you sure? [y/n]: "
    response = gets
    case response
    when /^[yY]/ then 
      return true
    else
      return false
    end
  end
end


s = 'hello'
s[0] = ?H
puts s


sure = are_you_sure?
if sure
  p "nice"
end

