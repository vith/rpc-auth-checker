# encoding: utf-8

require "xmlrpc/server"

s = XMLRPC::Server.new(10030)

SuccessfulAuth = Struct.new(:result, :account)
UnsuccessfulAuth = Struct.new(:error)

s.add_handler("checkAuthentication") do |user, pass|
	if user == 'foo' && pass == 'bar'
		SuccessfulAuth.new("Success", "someaccount")
	else
		UnsuccessfulAuth.new("Invalid password")
	end
end

s.set_default_handler do |name, *args|
	puts args
	raise XMLRPC::FaultException.new(-99, "Method #{name} missing" +
									 " or wrong number of parameters!")
end

s.serve
