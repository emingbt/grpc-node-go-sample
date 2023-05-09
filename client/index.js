const client = require('../client')

const path = require('path')
const express = require('express')
const bodyParser = require('body-parser')

const app = express()

app.use(bodyParser.json())
app.use(bodyParser.urlencoded({ extended: false }))

app.get('/', (req, res) => {
  const { month } = req.body

  client.calculateBendingType({ month: month }, (err, data) => {
    if (err) {
      return res.status(500).json({ message: err.message })
    }
    return res.status(200).json({ message: data.message })
  })
})

app.listen(3000, () => {
  console.log('Server is running on port 3000')
})