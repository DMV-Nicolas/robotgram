import { Navigate } from 'react-router-dom'
import { useToken } from '../hooks/useToken'

export function IsLoggedMiddleware({ children }: { children?: React.ReactNode }) {
  const { accessToken, refreshToken } = useToken()
  if (accessToken === '' || refreshToken === '') {
    return <Navigate to="/login" />
  } else {
    return children
  }
}
