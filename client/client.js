const PROTO_PATH = '../proto/ficsithub.proto'

const grpc = require('@grpc/grpc-js')
const protoLoader = require('@grpc/proto-loader')

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true
})

const FicsitHub = grpc.loadPackageDefinition(packageDefinition).FicsitHub
const client = new FicsitHub('localhost:50051', grpc.credentials.createInsecure())

module.exports = client