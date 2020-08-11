import axios from 'axios'

const http = axios.create({
  baseURL: '/api',
})

export abstract class BaseService {
  abstract path: string
}

export default http
