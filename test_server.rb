# encoding: utf-8

require "xmlrpc/server"

s = XMLRPC::Server.new(10030)

s.add_handler("checkAuthentication") do |user, pass|
	user == 'foo' && pass == 'bar'
end

s.set_default_handler do |name, *args|
	puts args
	raise XMLRPC::FaultException.new(-99, "Method #{name} missing" +
									 " or wrong number of parameters!")
end

s.serve
