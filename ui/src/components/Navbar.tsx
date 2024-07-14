import { Link } from 'react-router-dom'
import Clock from './Clock'
import Volume from './Volume'
import HomeSvg from '../images/home.svg'

export default function Navbar() {
  return (
    <div className="bg-background-inverted text-foreground-inverted shadow-md p-4 flex items-center gap-x-4">
      <div className="grow">
        <Link to="/">
          <img src={HomeSvg} className="w-12" />
        </Link>
      </div>
      <Clock />
      <Volume />
    </div>
  )
}
