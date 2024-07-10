import { useEffect, useState } from 'react'

function zeroPad(v: number): string {
  return String(v).padStart(2, '0')
}

export default function Clock() {

  const [time, setTime] = useState('')

  useEffect(() => {
    const updateTime = () => {
      const d = new Date()
      setTime(
        `${zeroPad(d.getHours())}:${zeroPad(d.getMinutes())}`
      )
    }
    updateTime()
    const id = setInterval(updateTime, 60 * 1000)
    return () => {
      clearInterval(id)
    }
  }, [])

  return (
    <div className="text-xl">
      {time}
    </div>
  )
}
