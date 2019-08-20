def wrap(wrappers, &block)
  yield
end

def with_logging(&block)
  puts "[log] starting block"
  yield
  puts "[log] ending block"
end

def with_error_handling(&block)
  begin
    yield
  rescue => e
    puts "[error] something bad happened: #{e}"
  end
end

def main
  wrap(with_error_handling, with_logging) do 
    puts "in body"
  end
end

main
