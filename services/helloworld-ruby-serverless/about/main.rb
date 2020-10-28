require 'sinatra'

set :bind, '0.0.0.0'
set :port, ENV['PT_CONTAINER_PORT'].to_i

get '/' do
  "Hello World - about page ruby\n"
end

