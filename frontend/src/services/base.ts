import axios from 'axios'

const http = axios.create({
  baseURL: 'localhost:8080/',
})

export abstract class BaseService {
  abstract path: string
}

export default http
