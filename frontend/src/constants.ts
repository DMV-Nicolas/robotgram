import { type PostType, type UserType } from './types'

export const DEFAULT_USER: UserType = {
  avatar: 'https://static.vecteezy.com/system/resources/thumbnails/009/292/244/small/default-avatar-icon-of-social-media-user-vector.jpg',
  createdAt: '',
  description: '',
  email: 'defaultuser@gmail.com',
  fullName: 'Default User',
  gender: 'male',
  id: '',
  username: 'defaultuser'
}

export const DEFAULT_POST: PostType = {
  id: '',
  userID: '',
  images: ['https://cdn2.iconfinder.com/data/icons/admin-tools-2/25/image2-512.png'],
  description: '',
  createdAt: ''
}
