import { type UserType } from '../types'

interface ProfileProps {
  user: UserType
}

export function Profile({ user }: ProfileProps) {
  return <h1>{user.gender}</h1>
}
