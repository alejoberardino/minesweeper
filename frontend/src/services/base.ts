import axios from 'axios'

const http = axios.create({
  baseURL: process.env.VUE_APP_API_BASE_URL,
})

export abstract class BaseService {
  abstract path: string
}

export default http
