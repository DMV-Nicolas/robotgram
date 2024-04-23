import { useState } from 'react'

export function useTransform({ transformModel }: { transformModel: ({ content }: { content: string }) => Promise<void> }) {
  const [transform, setTransform] = useState(() => transformModel)

  const updateTransform = ({ newTransform }: { newTransform: ({ content }: { content: string }) => Promise<void> }) => {
    setTransform(() => newTransform)
  }

  return { transform, updateTransform }
}
