'use client'

import Navbar from "./components/navbar";
import YoutubeCard from "./components/YoutubeCard";
import { useEffect, useState } from "react";
import { Video } from "@/lib/types";
import backend from "@/lib/axios";
import { toast } from "react-toastify";

export default function Home() {
  const [videoData, setVideoData] = useState<Video[]>([]);
  const [search, setSearch] = useState<string>("");
  const [page, setPage] = useState<number>(1);
  const [limit, setLimit] = useState<number>(15);

  const fetchData = async (search: string, page: number, limit: number) => {
    const id = toast.loading("Please wait...")

    try {
      const response = await backend.get(`/videos?title=${search}&page=${page}&limit=${limit}`)
      setVideoData(response.data.videos)
      toast.update(id, { render: "Videos fetched!", type: "success", isLoading: false });
      
    } catch (error: any) {
      toast.update(id, { render: "An error occured", type: "error", isLoading: false });
      console.log(error.response.data.message)
    }
  }

  useEffect(() => {
    fetchData(search, page, limit)
  }, [search, page, limit])

  return (
    <main>
      <Navbar setSearch={setSearch} />

      <section className="flex flex-row items-start justify-between flex-wrap w-[80%] mx-auto gap-8">
        {videoData.map((video, index) => {
          return (
            <YoutubeCard
              key={index}
              Thumbnail={video.thumbnail_url}
              VideoID={video.video_id}
              Title={video.title}
              PublishedAt={video.published_at}
              Description={video.description}
            />
          );
        })}
      </section>

      <section className="flex flex-row items-center justify-center py-8">
        <button
          className="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded-l"
          onClick={() => setPage(page - 1)} disabled={page === 1}>Previous</button>
        <button
          className="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded-r"
          onClick={() => setPage(page + 1)}>Next</button>
      </section>

    </main>
  );
}
