import { useEffect, useState } from 'react'
import Button from '../components/Button'
import VideoSvg from '../images/video.svg'

export default function Videos() {

  const [curVideo, setCurVideo] = useState<string | null>(null)
  const [videos, setVideos] = useState<string[]>([])

  useEffect(() => {
    fetch('/api/videos')
      .then(v => v.json())
      .then(v => setVideos(v))
  }, [])

  return (
    <>
      {curVideo !== null &&
        <video
          autoPlay
          className="absolute w-full h-full left-0 top-0 z-10"
          onEnded={() => setCurVideo(null)}
        >
          <source src={`/fs/videos/${curVideo}`} />
        </video>
      }
      <div className="flex flex-col gap-y-2">
        {videos.map((v, i) => (
          <Button
            key={i}
            onClick={() => setCurVideo(v)}
          >
            <img src={VideoSvg} className="w-8" />
            {v}
          </Button>
        ))}
      </div>
    </>
  )
}
