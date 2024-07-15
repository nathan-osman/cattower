import { Link } from 'react-router-dom'
import Clock from './Clock'
import Volume from './Volume'
import HomeSvg from '../images/home.svg'

export default function Navbar() {
  return (
    <div className="bg-background-inverted text-foreground-inverted shadow-md p-4 flex items-center gap-x-4">
      <Link to="/">
        <img src={HomeSvg} className="w-12" />
      </Link>
      <div className="grow" />
      <Clock />
      <Volume />
    </div>
  )
}
