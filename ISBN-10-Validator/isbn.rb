=begin
  This is the ISBN validator, written in Ruby instead of C, based on what I've
  learned from Codecademy and other resources. I'll add the functionality
  of checking ISBN-13 numbers too.
=end

def check_10(s)
  sum = 0
  i = 0
  cArray = s.chars
  #loop to sum the digits
  while i < 10
    sum += (10-i)*cArray[i].to_i
    puts "Sum incremented to: #{sum}"
    i+=1
  end
  if sum == 0
    puts "Invalid ISBN"
  elsif (sum % 11) == 0
    puts "Valid ISBN-10"
  else
    puts "Invalid ISBN-10"
  end
end

puts "Enter an ISBN number to be checked:"
input = gets.chomp

if (input.length == 10)
  check_10(input)
else
  puts "Invalid input length"
end
