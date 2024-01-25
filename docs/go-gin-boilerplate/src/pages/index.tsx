import useDocusaurusContext from "@docusaurus/useDocusaurusContext";
import Layout from "@theme/Layout";

const features = [
  {
    title: "快速开发",
    description: "使用 Gin 框架和相关工具，加速项目的开发和迭代过程。",
  },
  {
    title: "简单易用",
    description:
      "遵循 project-layout 规范, 提供清晰简单的代码结构，使新手也能轻松上手。",
  },
  {
    title: "先进的 CLI 体验",
    description: "使用 Cobra 打造现代命令行工具，简化项目管理和操作。",
  },
  {
    title: "热重载",
    description: "使用 Air 工具，支持热重载，提高开发效率。",
  },
  {
    title: "一体化日志系统",
    description: "集成 Logrus 和 Lumberjack, 实现全方位的日志记录和管理。",
  },
  {
    title: "数据库支持",
    description: "集成 Gorm, 支持主流数据库，如 MySQL、PostgreSQL 等。",
  },
  {
    title: "灵活的中间件",
    description: "整合常用中间件，轻松实现日志、认证、跨域、限流等功能。",
  },
  {
    title: "API 文档",
    description: "使用 Gin-Swagger 生成 API 文档，方便查看和调试接口。",
  },
];

export default function Home(): JSX.Element {
  const { siteConfig } = useDocusaurusContext();
  return (
    <Layout
      title={`Hello from ${siteConfig.title}`}
      description="A boilerplate for Go Gin Web Framework."
    >
      <div className=" py-20 bg-gradient-to-tr from-indigo-100 to-white dark:bg-none">
        <div className=" px-4 sm:px-10">
          <div className=" text-4xl sm:text-5xl font-bold text-center">
            Develop. Fast. First.
          </div>

          <div className=" mt-10 mb-10 text-xl sm:text-2xl text-center">
            <div className="text-3xl sm:text-4xl font-bold mb-8">
              {siteConfig.title}
              <img src="img/golang.png" alt="" className="h-10 ml-2" />
            </div>
            <div className=" sm:max-w-[50%] text-center mx-auto">
              A development boilerplate based on the Gin framework, aimed at
              helping developers quickly build and develop web applications.
            </div>
            <div className=" mt-4 flex items-center gap-4 mx-auto justify-center">
              <img
                onClick={() =>
                  window.open(
                    "https://github.com/sanjayheaven/go-gin-boilerplate"
                  )
                }
                src="https://img.shields.io/github/stars/sanjayheaven/go-gin-boilerplate"
                alt=""
                className=" h-6 cursor-pointer"
              />

              <img
                onClick={() =>
                  window.open(
                    "https://github.com/sanjayheaven/go-gin-boilerplate"
                  )
                }
                src="https://github.com/sanjayheaven/go-gin-boilerplate/workflows/Go/badge.svg"
                alt=""
                className=" h-6 cursor-pointer"
              />
            </div>
            {/* {siteConfig.tagline} */}
          </div>
        </div>

        <div className=" flex flex-wrap items-center justify-center gap-6 mx-auto px-2 sm:px-10">
          <div
            className=" box-border sm:w-[200px] sm:max-w-[200px] hover:shadow-xl cursor-pointer py-3 sm:py-4 px-4 sm:px-6 font-bold text-2xl 
          transition-all duration-300 bg-white rounded-md text-black no-underline
          min-w-[170px] sm:min-w-max text-center"
            onClick={() => (window.location.href = "docs/intro")}
          >
            Get Started
          </div>
          <div
            className=" box-border sm:w-[200px] sm:max-w-[200px] hover:shadow-xl cursor-pointer py-3 sm:py-4 px-4 sm:px-10 font-bold text-2xl 
          transition-all duration-300 bg-white rounded-md text-black no-underline
          min-w-[170px] sm:min-w-max text-center"
            onClick={() =>
              window.open("https://github.com/sanjayheaven/go-gin-boilerplate")
            }
          >
            GitHub
          </div>
        </div>
      </div>

      <main>
        {/* quick start */}
        <div className=" my-20">
          <div className="text-3xl sm:text-4xl font-bold text-center mb-10">
            Quick Start
          </div>

          <div className=" sm:max-w-[50%] mx-auto px-4">
            {/* <div className="flex items-center gap-10 mb-6">
              <div className=" cursor-pointer bg-white border rounded-md border-solid w-[150px] text-center shadow-md px-5 py-2">
                CLI
              </div>
              <div className=" cursor-pointer bg-white border rounded-md border-solid w-[150px] text-center shadow-md px-5 py-2">
                Source Code
              </div>
            </div> */}

            <div>
              <div className=" font-bold text-2xl mb-4">Installations</div>
              <div className=" w-full flex items-center justify-center mx-auto gap-6">
                <code
                  lang="sh"
                  className=" w-full p-4 sm:p-8 bg-[#2b303b] text-white"
                >
                  <div>
                    git clone
                    https://github.com/sanjayheaven/go-gin-boilerplate.git
                  </div>
                  <div>cd go-gin-boilerplate </div>
                  <div>go mod download</div>
                </code>
              </div>
            </div>
            {/* config */}
            <div className=" mt-8">
              <div className=" font-bold text-2xl mb-4">Config Setup</div>
              <div className=" w-full flex items-center justify-center mx-auto gap-6">
                <code
                  lang="sh"
                  className=" w-full p-4 sm:p-8 bg-[#2b303b] text-white"
                >
                  <div>cp configs/config.example.yaml configs/config.yaml</div>
                  <div>vi configs/config.yaml</div>
                </code>
              </div>
            </div>
            {/* run */}
            <div className=" mt-8">
              <div className=" font-bold text-2xl mb-4">Run Project</div>
              <div className=" w-full flex items-center justify-center mx-auto gap-6">
                <code
                  lang="sh"
                  className=" w-full p-4 sm:p-8 bg-[#2b303b] text-white"
                >
                  <div>go run main.go</div>
                </code>
              </div>
            </div>
          </div>
        </div>

        {/* why  */}
        <div className=" my-20">
          <div className="text-3xl sm:text-4xl font-bold text-center mb-10">
            Why use this boilerplate?
          </div>

          <div className=" md:grid grid-cols-3 justify-items-center w-max mx-auto gap-6">
            {features.map((item) => {
              return (
                <div
                  key={item.title}
                  className=" hover:scale-110 max-w-max cursor-pointer shadow-lg rounded-lg w-[250px] p-8 transition-all duration-300"
                >
                  <div className="font-bold ">{item.title}</div>
                  <div>{item.description}</div>
                </div>
              );
            })}
          </div>
        </div>
        {/*  */}
      </main>

      {/* footer */}

      <div className=" my-20 text-center text-4xl font-bold">
        Develop. Fast. First.
      </div>
    </Layout>
  );
}
