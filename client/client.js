const PROTO_PATH = '../bending/bending.proto'

const grpc = require('@grpc/grpc-js')
const protoLoader = require('@grpc/proto-loader')

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true
})

const BendingService = grpc.loadPackageDefinition(packageDefinition).bending
const client = new BendingService.Bending('localhost:50051', grpc.credentials.createInsecure())
module.exports = client