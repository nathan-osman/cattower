import { useEffect, useState } from 'react'

type Event = {
  timestamp: string
  motion: boolean
}

export default function MotionLog() {

  const [log, setLog] = useState<Event[]>([])

  useEffect(() => {
    fetch('/api/motion/log')
      .then(v => v.json())
      .then(v => setLog(v))
  }, [])

  return (
    <div className="h-full flex flex-col gap-y-2">
      <div>Motion Log:</div>
      <div className="border border-foreground grow p-2 font-mono overflow-auto">
        {
          log.map((v, i) => (
            <div key={i}>
              {new Date(v.timestamp).toString()} - {v.motion ? "active" : "inactive"}
            </div>
          ))
        }
      </div>
    </div>
  )
}
