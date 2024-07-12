import { Outlet } from 'react-router-dom'
import Navbar from './Navbar'

export default function App() {
  return (
    <div className="flex flex-col h-full overflow-hidden">
      <Navbar />
      <div className="p-4 grow overflow-auto">
        <Outlet />
      </div>
    </div>
  )
}
