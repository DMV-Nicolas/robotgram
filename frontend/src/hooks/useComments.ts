import { useEffect, useState } from 'react'
import { type CommentType } from '../types'

export function useComments({ targetID }: { targetID: string }) {
  const [comments, setComments] = useState<CommentType[]>([])

  const mockComments: CommentType[] = [
    {
      id: 'qwerty123456789id1',
      targetID: 'qwerty123456789target1',
      userID: 'qwerty123456789user1',
      content: 'Hello world!',
      createdAt: '2024-02-17T15:22:34.517Z'
    },
    {
      id: 'qwerty123456789id2',
      targetID: 'qwerty123456789target2',
      userID: 'qwerty123456789user2',
      content: 'Hola perras!!',
      createdAt: '2024-02-17T15:22:34.517Z'
    },
    {
      id: 'qwerty123456789id3',
      targetID: 'qwerty123456789target3',
      userID: 'qwerty123456789user3',
      content: 'Hello world!',
      createdAt: '2024-02-17T15:22:34.517Z'
    },
    {
      id: 'qwerty123456789id4',
      targetID: 'qwerty123456789target4',
      userID: 'qwerty123456789user4',
      content: 'Oi manito puto! Oi manito puto!!, lorem ipsum Oi manito puto! Oi manito puto!!, lorem ipsum Oi manito puto! Oi manito puto!!, lorem ipsum Oi manito puto! Oi manito puto!!, lorem ipsum door',
      createdAt: '2024-02-17T15:22:34.517Z'
    },
    {
      id: 'qwerty123456789id5',
      targetID: 'qwerty123456789target5',
      userID: 'qwerty123456789user5',
      content: 'ARROZ CON PANELA PERRO',
      createdAt: '2024-02-17T15:22:34.517Z'
    },
    {
      id: 'qwerty123456789id1',
      targetID: 'qwerty123456789target1',
      userID: 'qwerty123456789user1',
      content: 'Hello world!',
      createdAt: '2024-02-17T15:22:34.517Z'
    },
    {
      id: 'qwerty123456789id2',
      targetID: 'qwerty123456789target2',
      userID: 'qwerty123456789user2',
      content: 'Hola perras!!',
      createdAt: '2024-02-17T15:22:34.517Z'
    },
    {
      id: 'qwerty123456789id3',
      targetID: 'qwerty123456789target3',
      userID: 'qwerty123456789user3',
      content: 'Hello world!',
      createdAt: '2024-02-17T15:22:34.517Z'
    },
    {
      id: 'qwerty123456789id4',
      targetID: 'qwerty123456789target4',
      userID: 'qwerty123456789user4',
      content: 'Oi manito puto! Oi manito puto!!, lorem ipsum Oi manito puto! Oi manito puto!!, lorem ipsum Oi manito puto! Oi manito puto!!, lorem ipsum Oi manito puto! Oi manito puto!!, lorem ipsum door',
      createdAt: '2024-02-17T15:22:34.517Z'
    },
    {
      id: 'qwerty123456789id5',
      targetID: 'qwerty123456789target5',
      userID: 'qwerty123456789user5',
      content: 'ARROZ CON PANELA PERRO',
      createdAt: '2024-02-17T15:22:34.517Z'
    },
    {
      id: 'qwerty123456789id1',
      targetID: 'qwerty123456789target1',
      userID: 'qwerty123456789user1',
      content: 'Hello world!',
      createdAt: '2024-02-17T15:22:34.517Z'
    },
    {
      id: 'qwerty123456789id2',
      targetID: 'qwerty123456789target2',
      userID: 'qwerty123456789user2',
      content: 'Hola perras!!',
      createdAt: '2024-02-17T15:22:34.517Z'
    },
    {
      id: 'qwerty123456789id3',
      targetID: 'qwerty123456789target3',
      userID: 'qwerty123456789user3',
      content: 'Hello world!',
      createdAt: '2024-02-17T15:22:34.517Z'
    },
    {
      id: 'qwerty123456789id4',
      targetID: 'qwerty123456789target4',
      userID: 'qwerty123456789user4',
      content: 'Oi manito puto! Oi manito puto!!, lorem ipsum Oi manito puto! Oi manito puto!!, lorem ipsum Oi manito puto! Oi manito puto!!, lorem ipsum Oi manito puto! Oi manito puto!!, lorem ipsum door',
      createdAt: '2024-02-17T15:22:34.517Z'
    },
    {
      id: 'qwerty123456789id5',
      targetID: 'qwerty123456789target5',
      userID: 'qwerty123456789user5',
      content: 'ARROZ CON PANELA PERRO',
      createdAt: '2024-02-17T15:22:34.517Z'
    }
  ]

  useEffect(() => {
    setComments(mockComments)
  }, [])

  return { comments }
}
