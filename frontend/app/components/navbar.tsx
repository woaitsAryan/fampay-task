'use client'

import Image from "next/image"

interface NavbarProps {
    setSearch: (search: string) => void
}

const Navbar = (props: NavbarProps) => {

    return (
        <header className="text-gray-600 body-font bg-slate-200 sticky top-0 mb-4">
            <div className="container mx-auto flex flex-wrap p-5 flex-col md:flex-row items-center">
                <a className="flex title-font font-medium items-center text-gray-900 mb-4 md:mb-0">
                    <Image src="/fampay.png" alt="logo" width={50} height={50} />
                    <span className="ml-3 text-xl">Fampay Task</span>
                </a>
                <div className="md:ml-auto md:mr-auto flex flex-wrap items-center text-base justify-center gap-4">
                    <input
                        type="text"
                        placeholder="Search"
                        className="shadow appearance-none border rounded w-96 py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                        onChange={(e) => props.setSearch(e.target.value)}
                    />
                    <button
                        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                    >
                        Search
                    </button>
                </div>
                {/* <button className="inline-flex items-center bg-gray-100 border-0 py-1 px-3 focus:outline-none hover:bg-gray-200 rounded text-base mt-4 md:mt-0">Button
                    <svg fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" className="w-4 h-4 ml-1" viewBox="0 0 24 24">
                        <path d="M5 12h14M12 5l7 7-7 7"></path>
                    </svg>
                </button> */}
            </div>
        </header>
    )
}

export default Navbar