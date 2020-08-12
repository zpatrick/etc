# https://medium.com/@shvetsovdm/playing-with-ruby-threads-and-queues-52beb6e8613c
class WorkerPool
  def initialize(size: 5)
    @queue = Queue.new
    @workers = Array.new(size) { Worker.new(@queue) }
  end

  def start
    @workers.each(&:start)
  end

  def stop
    @queue.close
    @workers.each(&:stop)
  end

  def perform(*payloads, &block)
    payloads.each { |payload| @queue.push([block, payload]) }
    while !@queue.empty?
  end
end

class Worker
  def initialize(queue)
    @queue = queue
  end

  def start
    @thread = Thread.new do
      while !queue.closed? || !queue.empty?
        proc, payload = queue.pop(false)
        proc.call(payload)
      end
    end
  end
end


class MessageService
  def bulk_get(namespace, mailbox, *message_ids)
    messages = []
    mutex = mutex.
    @pool.enqueue(message_ids) do |message_id|
      @message_reader.read_message(namespace, mailbox, message_id)
    end
  end
end
