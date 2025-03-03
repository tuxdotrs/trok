import { createSignal } from "solid-js";
import {
  ChatBubbleLeftIcon,
  ClipboardDocumentIcon,
  ClipboardDocumentCheckIcon,
} from "./icons";

function App() {
  const install = `curl -fsSL https://trok.cloud/install.sh | sh`;
  const [copied, setCopied] = createSignal(false);

  const handleCopy = () => {
    navigator.clipboard.writeText(install);
    setCopied(true);
    setTimeout(() => {
      setCopied(false);
    }, 1300);
  };

  return (
    <div class="md:flex md:h-screen md:w-screen py-8 justify-center items-center antialiased">
      <div class="mx-auto lg:max-w-5xl">
        <div class="mx-5 md:mx-0 flex flex-col justify-center items-center">
          <h1 class="text-7xl text-zinc-700 font-koulen">TROK</h1>
          <h3 class="text-lg text-zinc-500 font-koulen">
            EXPOSE. SHARE. ACCESS.
          </h3>
        </div>

        <div class="mx-5 md:mx-0 h-10 border-gray-200 border-l border-r" />

        <div class="mx-5 md:mx-0 border-t border-gray-200 border-l border-r">
          <div class="md:flex">
            <div class="px-5 py-10 border-b border-gray-200 md:border-r md:w-5/12">
              <h2 class="text-4xl text-zinc-700 font-koulen">
                Accessing your
                <br />
                local service should
                <br />
                be simple
              </h2>
              <p class="mt-5 text-md/10 text-gray-500">
                It&apos;s like&nbsp;
                <span class="font-mono bg-zinc-100 px-1 py-0.5 rounded">
                  ngrok
                </span>
                &nbsp;but simpler. Just log in, run the command, and get a
                shareable URL to access your local service from anywhere. No
                complex setup—just instant, secure tunneling to the internet.
              </p>
            </div>

            <div class="px-5 py-10 border-b border-gray-200 flex-1 md:w-7/12">
              <h3 class="text-3xl text-zinc-700 font-koulen">Install</h3>
              <p class="mt-1 text-sm/6 text-gray-500">
                Install trok to your machine using shell:
              </p>
              <button
                class="text-xs md:text rounded bg-zinc-100 p-2 mt-2 flex cursor-pointer"
                onclick={handleCopy}
              >
                <div class="grow mr-10 text-zinc-600 font-mono">
                  $ {install}
                </div>
                {copied() ? (
                  <span class="w-10 mr-1">copied!</span>
                ) : (
                  <span class="w-10 mr-1"></span>
                )}
                {copied() ? (
                  <ClipboardDocumentCheckIcon class="h-5 w-5 text-zinc-500" />
                ) : (
                  <ClipboardDocumentIcon class="h-5 w-5 text-zinc-500" />
                )}
              </button>
              <div class="mt-2 text-xs/6 text-gray-500">
                <p>
                  This will download the trok binary to the path you ran the
                  script from.
                  <br />
                  Run it with&nbsp;
                  <span class="font-mono text-sm/6 rounded bg-zinc-100 p-1 inline">
                    ./trok
                  </span>
                  &nbsp;on any system
                </p>
              </div>
            </div>
          </div>
        </div>

        <div class="mx-5 md:mx-0 h-5 border-l border-gray-200 border-r" />
        <div class="mx-5 md:mx-0 border-t border-gray-200 border-l border-r">
          <div class="px-5 pt-10">
            <h3 class="font-koulen text-zinc-600 text-xl">USING TROK:</h3>
          </div>
          <div class="md:flex sm:gap-5 border-gray-200 border-b">
            <div class="md:w-1/3 p-5 pb-10">
              <div class="rounded border border-zinc-300 bg-zinc-100 p-2 pb-8 h-40 mb-5">
                <div class="flex gap-1 items-center">
                  <div class="border border-zinc-400 w-3 h-3 rounded-full" />
                  <div class="border border-zinc-400 w-3 h-3 rounded-full" />
                  <div class="border border-zinc-400 w-3 h-3 rounded-full" />
                </div>
                <p class="font-mono text-zinc-700 font-bold text-sm pt-4">
                  {">"} $ trok client -a :3000
                </p>
                <p class="font-mono text-zinc-400 text-sm">
                  started Trok client on trok.cloud
                </p>
                <p class="font-mono text-zinc-400 text-sm">
                  [CMD] EHLO [ARG] trok.cloud
                </p>
              </div>
              <h3 class="font-koulen text-zinc-700 text-3xl">1. Setup</h3>
              <p class="mt-1 text-sm/6 text-gray-500">
                provide the local port you want to share, and it&apos;ll
                generate a URL.
              </p>
            </div>
            <div class="md:w-1/3 p-5 pb-10">
              <div class="relative p-2 flex h-40 mb-5">
                <ChatBubbleLeftIcon class="w-6 h-6 mr-2" />
                <div class="rounded-tr-lg rounded-br-lg rounded-bl-lg border border-gray-200 bg-zinc-50 p-2 pb-10 mb-2 h-20">
                  hey here take a look at my project...
                </div>
              </div>
              <h3 class="font-koulen text-zinc-700 text-3xl">2. Share</h3>
              <p class="mt-1 text-sm/6 text-gray-500">
                share the URL with your friends.
              </p>
            </div>
            <div class="md:w-1/3 p-5 pb-10">
              <div class="rounded border border-zinc-300 bg-zinc-100 p-2 pb-8 h-40 mb-5">
                <div class="flex gap-1 items-center">
                  <div class="flex gap-1">
                    <div class="border border-zinc-400 w-3 h-3 rounded-full" />
                    <div class="border border-zinc-400 w-3 h-3 rounded-full" />
                    <div class="border border-zinc-400 w-3 h-3 rounded-full" />
                  </div>
                  <div class="border flex items-center border-zinc-400 w-full h-4 rounded-sm">
                    <span class="text-xs text-zinc-400 px-1 -mt-[1px]">
                      https://awesome-project.trok.cloud
                    </span>
                  </div>
                </div>
                <div class="flex justify-center items-center h-full font-mono text-zinc-400 text-sm">
                  🎉 your awesome project 🎉
                </div>
              </div>
              <h3 class="font-koulen text-zinc-700 text-3xl">3. Browse</h3>
              <p class="mt-1 text-sm/6 text-gray-500">
                access your local service from anywhere.
              </p>
            </div>
          </div>
        </div>
        <div class="mx-5 md:mx-0 h-10 border-gray-200 border-l border-r" />
      </div>
    </div>
  );
}

export default App;
