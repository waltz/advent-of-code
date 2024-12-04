file = File.open('./input')

counter = 0
frequency_occurrences = {}
duplicate_frequency = nil

while duplicate_frequency.nil? do
  puts "LOOKING AT LINES"
  file.seek(0)
  file.each_line do |line|
    match = line.match(/^(\+|-*)(\d+)$/)

    operator = match[1].to_s
    amount = match[2].to_i

    if operator == '+'
      counter = counter + amount
    elsif operator == '-'
      counter = counter - amount
    else
      puts "unknown operator: #{operator}"
    end

    if frequency_occurrences[counter].nil?
      frequency_occurrences[counter] = 1
    else
      frequency_occurrences[counter] = frequency_occurrences[counter] + 1
    end


    puts "current frequency is: #{counter}, we have seen this frequncy #{frequency_occurrences[counter]} times"
  
    if frequency_occurrences[counter] == 2
      duplicate_frequency = counter
      break
    end
  end
end

puts "Duplicate frequency is: #{duplicate_frequency}"
