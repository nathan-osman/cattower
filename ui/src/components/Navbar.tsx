import Clock from './Clock'
import HomeSvg from '../images/home.svg'
import { Link } from 'react-router-dom'

export default function Navbar() {
  return (
    <div className="bg-background-inverted text-foreground-inverted shadow-md p-4 flex items-center">
      <div className="grow">
        <Link to="/">
          <img src={HomeSvg} className="w-8" />
        </Link>
      </div>
      <Clock />
    </div>
  )
}
