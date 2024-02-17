import { usePosts } from '../hooks/usePosts'
import { type PostType, type UserType } from '../types'
import { Options } from './Icons'
import './Profile.css'

interface ProfileHeaderProps {
  username: string
  fullName: string
  email: string
  avatar: string
  description: string
  gender: string
}

function ProfileHeader({ username, fullName, email, avatar, description, gender }: ProfileHeaderProps) {
  // const genderIcon = gender === 'male' ? <Male /> : <Female />
  return (
    <div className='profileHeader'>
      <div className='profileHeader__left'>
        <img className='profileHeader__avatar' src={avatar} alt="" />
      </div>
      <div className='profileHeader__right'>
        <div className='profileHeader__info'>
          <h2 className='profileHeader__username'>{username}</h2>
          <button className='profileHeader__button'>Follow</button>
          <button className='profileHeader__buttonIcon'>
            <Options />
          </button>
        </div>
        <div className='profileHeader__info profileHeader__info--column'>
          <span className='profileHeader__fullname'>{fullName} - {email}</span>
          <p>{description}</p>
        </div>
      </div>
    </div>
  )
}

interface ProfileBodyProps {
  posts: PostType[]
}

function ProfileBody({ posts }: ProfileBodyProps) {
  return (
    <div className='profileBody'>
      <ul className='profileBody__ul'>
        {posts.map((post) => (
          <li className='profileBody__li' key={post.id}>
            <img className='profileBody__postImage' src={post.images[0]} alt="" />
          </li>
        ))}
      </ul>
    </div>
  )
}

interface ProfileProps {
  user: UserType
}

export function Profile({ user }: ProfileProps) {
  const { posts } = usePosts({ userID: user.id })
  return (
    <div className='profileContainer'>
      <div className='profile'>
        <ProfileHeader
          username={user.username}
          fullName={user.fullName}
          email={user.email}
          avatar={user.avatar}
          description={user.description}
          gender={user.gender}
        />
        <hr />
        <ProfileBody
          posts={posts}
        />
      </div>
    </div>
  )
}
