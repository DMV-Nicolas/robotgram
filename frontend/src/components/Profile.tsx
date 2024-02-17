import { type UserType } from '../types'
import { Female, Male, Options } from './Icons'
import './Profile.css'

interface ProfileHeaderProps {
  username: string
  fullName: string
  email: string
  avatar: string
  description: string
  gender: string
}

export function ProfileHeader({ username, fullName, email, avatar, description, gender }: ProfileHeaderProps) {
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
          <span>{fullName} - {email}</span>
          <p>{description}</p>
        </div>
      </div>
    </div>
  )
}

interface ProfileProps {
  user: UserType
}

export function Profile({ user }: ProfileProps) {
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
      </div>
    </div>
  )
}
