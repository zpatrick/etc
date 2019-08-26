class Person
  attr_reader :email
  def initialize(name, age)
    @name, @age = name, age
  end
  def name=(n)
    @name=n
  end
  def age=(a)
    @age=a
  end
  def to_s
    return "#{@name} (age: #{@age})"
  end
end

zack = Person.new("Zack", 27)
zack.email = "foo"
puts zack
