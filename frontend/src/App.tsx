import { useState } from 'react'

import { ProfileForm } from '@/components/auth/user-auth-form'

import img from './assets/wall.jpg'
import ico from './assets/ico.png'

function App() {
  return (
    <>
      <div className='container relative h-screen w-screen md:flex-row flex-col m-auto justify-between max-w-none flex px-0'>
        <div className="grow-[8] relative h-full md:w-fit w-full flex-col bg-muted p-8 text-white dark:border-r flex">
          <div className="absolute bg-center h-full inset-0 bg-cover" style={{ backgroundImage: `url(${img})` }} />
          <a
            href="https://ki365.github.io/"
            className="relative mb-5 z-20 flex items-center text-xl font-medium">
            <img src={ico} className="mr-2 h-7 w-7" />
            Ki365
          </a>
          <div className="relative z-25 mt-auto">
            <blockquote className="space-y-2">
              Photo by <a href="https://unsplash.com/@efekurnaz?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Efe Kurnaz</a> on <a href="https://unsplash.com/photos/multicolored-hallway-RnCPiXixooY?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Unsplash</a>
            </blockquote>
          </div>
        </div>
        <div className="grow"></div>
        <div className="overflow-visible m-auto">
          <div className="m-10 z-30 flex w-full flex-col justify-center space-y-6 w-[325px]">
            <div className="flex z-30 flex-col space-y-2 text-center !w-full">
              <h1 className="text-2xl z-30 font-semibold tracking-tight !w-full">
                Log in
              </h1>
              <p className="text-sm text-muted-foreground">
                Enter your credentials below to login to Ki365
              </p>
            </div>
            <ProfileForm />
            <p className="px-8 text-center text-sm text-muted-foreground">
              This project is made with ❤️ @ {" "}
              <a
                href="https://github.com/Ki365/Ki365"
                className="underline underline-offset-4 hover:text-primary"
              >
                Github Repo
              </a>
            </p>
            <p className="px-8 text-center text-sm text-muted-foreground">
              By continuing, you agree to this site's {" "}
              <a
                href="/terms"
                className="underline underline-offset-4 hover:text-primary"
              >
                Terms of Service
              </a>{" "}
              and{" "}
              <a
                href="/privacy"
                className="underline underline-offset-4 hover:text-primary"
              >
                Privacy Policy
              </a>
              .
            </p>
          </div>
        </div>
        <div className="grow"></div>
      </div>
    </>
  )
}

export default App
