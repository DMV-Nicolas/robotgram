import './Footer.css'

export function Footer({ previousLocation }: { previousLocation: string }) {
  return (
    <footer className='footer'>
      {previousLocation}
    </footer>
  )
}
