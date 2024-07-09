import { Outlet } from 'react-router-dom'
import Navbar from './Navbar'

export default function App() {
  return (
    <div>
      <Navbar />
      <div className="p-4">
        <Outlet />
      </div>
    </div>
  )
}
