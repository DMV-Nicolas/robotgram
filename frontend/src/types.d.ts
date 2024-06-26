export interface UserResponse {
  id: string
  username: string
  full_name: string
  email: string
  avatar: string
  description: string
  gender: string
  created_at: string
}

export interface UserType {
  id: string
  username: string
  fullName: string
  email: string
  avatar: string
  description: string
  gender: string
  createdAt: string
}

export interface UsersLoginResponse {
  session_id: string
  access_token: string
  access_token_expires_at: string
  refresh_token: string
  refresh_token_expires_at: string
}

export interface PostResponse {
  id: string
  user_id: string
  images: string[]
  description: string
  created_at: string
}

export interface PostType {
  id: string
  userID: string
  images: string[]
  description: string
  createdAt: string
}

export type ListPostsResponse = PostResponse[] | null

export interface CommentResponse {
  id: string
  user_id: string
  target_id: string
  content: string
  created_at: string
}

export interface CommentType {
  id: string
  userID: string
  targetID: string
  content: string
  createdAt: string
}

export type ListCommentsResponse = CommentResponse[] | null

export interface RefreshTokenResponse {
  access_token: string
  access_token_expires_at: string
}

export interface GetTokenDataResponse {
  id: string
  user_id: string
  issued_at: string
  expires_at: string
}

export interface TokenContextType {
  accessToken: string
  refreshToken: string
  updateAccessToken: (newToken: string) => void
  updateRefreshToken: (newToken: string) => void
  refreshAccessToken: () => Promise<Error | undefined>
}

export type LikesCountResponse = number

export type IsLikedResponse = boolean

export interface CreatedResponse {
  InsertedID: string
}
