import { useEffect, useState } from 'react'
import { matchPath, useLocation } from 'react-router-dom'

export function useModals({ modalPages }: { modalPages: string[] }) {
  const location = useLocation()
  const [previousLocation, setPreviousLocation] = useState('')

  const shouldUpdatePreviousLocation = () => {
    const f = modalPages.filter((path) => (
      matchPath(path, location.pathname) !== null
    ))
    return f.length !== 1
  }

  useEffect(() => {
    if (shouldUpdatePreviousLocation()) {
      setPreviousLocation(location.pathname)
    }
  }, [location.pathname])

  return { previousLocation }
}
