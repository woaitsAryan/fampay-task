import { formatTimeDifference } from "@/lib/time"
import Image from "next/image"
import Link from "next/link"


interface YoutubeCardProps {
    Thumbnail: string
    Title: string
    PublishedAt: string
    VideoID: string
    Description: string
}

const YoutubeCard = (props: YoutubeCardProps) => {
    return(
        <Link href={`https://www.youtube.com/watch?v=${props.VideoID}`} className = "w-[30%] aspect-[25/20]">
            <Image src = {props.Thumbnail} alt = {props.Description} width = {1280} height = {720} className="h-[75%] rounded-lg mb-2"/>
            <p className="text-lg font-medium truncate">{props.Title}</p>
            <p className="text-sm">{formatTimeDifference(props.PublishedAt)}</p>
        </Link>  
    )
}

export default YoutubeCard  