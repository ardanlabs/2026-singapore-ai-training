**Subject:** Ultimate Practical AI Workshop prep

Hi all,

Ahead of the workshop, please get your machine ready by installing a few tools
and pulling the model files. The downloads are large (~50 GB total), so please
do this on a good connection well before we start.

If your machine does not have a GPU with at least 8GB of VRAM, do not worry, you
can still be part of the class, I will provide you with a way to run these examples.

Everything we'll learn in the class is applicable to frontier models/providers
(OpenAI/Anthropic/Qwen/...) or local models.

**1. Prerequisites**

- Go 1.26 — https://go.dev/dl/
  (or `go install golang.org/dl/go1.26@latest && go1.26 download`)
- Homebrew — https://brew.sh (optional)
- Docker (give it at least 4 CPUs)

Homebrew is optional. You can install the dependencies required using other
package managers or manually.

**2. Install CLI tooling**

```shell
go install github.com/ardanlabs/kronk/cmd/kronk@latest
brew install mongosh mplayer pkgconf pgcli uv whisper-cpp
brew tap homebrew-ffmpeg/ffmpeg
brew install homebrew-ffmpeg/ffmpeg/ffmpeg --with-whisper-cpp
go install github.com/janpfeifer/gonb@latest
```

**3. Pull the Docker images**

```shell
docker pull pgvector/pgvector:pg18
docker pull mongodb/mongodb-atlas-local:8.2
docker pull quay.io/docling-project/docling-serve:v1.9.0
```

**4. Start Kronk (leave this terminal open)**

```shell
kronk server start
```

**5. Pull the models (in a new terminal)**

```shell
kronk model pull --local "unsloth/Qwen3-0.6B-Q8_0"
kronk model pull --local "unsloth/Qwen3-8B-GGUF/Qwen3-8B-Q8_0.gguf"
kronk model pull --local "ggml-org/embeddinggemma-300m-qat-q8_0-GGUF/embeddinggemma-300m-qat-Q8_0.gguf"
kronk model pull --local "ggml-org/Qwen2.5-VL-3B-Instruct-GGUF/Qwen2.5-VL-3B-Instruct-Q8_0.gguf"
kronk model pull --local "mradermacher/Qwen2-Audio-7B-GGUF/Qwen2-Audio-7B.Q8_0.gguf"
kronk model pull --local "unsloth/gpt-oss-20b-GGUF:Q8_0"
kronk model pull --local "bartowski/cerebras_Qwen3-Coder-REAP-25B-A3B-GGUF/cerebras_Qwen3-Coder-REAP-25B-A3B-Q8_0.gguf"
```

**Verify**

```shell
kronk model list --local
```

The repo itself will be shared at the start of the workshop.

E-mail me at florin.patan@ardanlabs.com for anything.

Thanks,
Florin
