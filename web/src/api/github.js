import axios from 'axios'

const service = axios.create()

export function Commits(page) {
  return service({
    url:
      'https://www.spiritlhl.net' +
      page,
    method: 'get'
  })
}

export function Members() {
  return service({
    url: 'https://www.spiritlhl.net',
    method: 'get'
  })
}
