file = File.open('./input')

counter = 0
frequency_occurrences = {}

file.each_line do |line|
  match = line.match(/^(\+|-*)(\d+)$/)

  operator = match[1].to_s
  amount = match[2].to_i

  puts "current operator is #{operator} and the amount is #{amount}"

  if operator == '+'
    counter = counter + amount
  elsif operator == '-'
    counter = counter - amount
  else
    puts "unknown operator: #{operator}"
  end

  puts "current count is: #{counter}"
end

puts "final count is: #{counter}"
