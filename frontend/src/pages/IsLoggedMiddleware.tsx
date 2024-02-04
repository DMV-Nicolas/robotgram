import { Navigate } from 'react-router-dom'
import { useToken } from '../hooks/useToken'

export function IsLoggedMiddleware({ children }: { children?: React.ReactNode }) {
  const { accessToken } = useToken()
  if (accessToken.current !== '') {
    return children
  } else {
    return <Navigate to="/login" />
  }
}
