class AmbiguousBases
	def initialize(num, base)
		@this_num = num
		@this_base = base
	end

	def findAlternatives()
		
	end

end

puts("Enter a number to be evaluated:")
num = gets.chomp()
puts("#{num} will be evaluated. Enter a base (e.g. 2)")
base = gets.chomp()

ab = AmbiguousBases.new(num, base)
ab.findAlternatives()
