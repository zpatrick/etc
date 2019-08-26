class EWSParsingError < StandardError
end

def with_retries(max_attempts=3, &block)
  (max_attempts - 1).times do |i|
    begin
      yield
    rescue EWSParsingError
      puts "Parsing error encountered, retrying..."
      next
    end
  end

  yield
end

#with_retries(3) do 
#  puts "Hello, World!"
#  raise EWSParsingError.new
# end


def attempt(max, &block) 
  (1..max).each do |i|
    yield i, max
  end
end

# attempt(Float::INFINITY) do |i|

ews_item_ids = attempt(5) do |i, max|
  begin
    # ews_folder.get_items(...)
  rescue EWSParseError
    puts "ParseError, trying again..." unless i == max
  end
end



# attempt((1...Inf)).do |i|
#  puts i
# end






















