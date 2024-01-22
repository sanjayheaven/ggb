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
      <div className=" py-20 bg-theme">
        <div className=" text-5xl font-bold text-center">
          {siteConfig.title}
          <img src="img/golang.png" alt="" className="h-10 ml-2" />
        </div>

        <div className=" mt-10 mb-16 text-2xl text-center">
          {siteConfig.tagline}
        </div>

        <div className=" flex items-center justify-center gap-6 ">
          <div
            className=" w-[200px] hover:shadow-md cursor-pointer py-4 px-10 font-bold text-2xl 
          transition-all duration-300 bg-white rounded-md text-black no-underline
          min-w-max text-center"
            onClick={() => (window.location.href = "docs/intro")}
          >
            Get Started
          </div>
          <div
            className=" w-[200px] hover:shadow-md cursor-pointer py-4 px-10 font-bold text-2xl 
          transition-all duration-300 bg-white rounded-md text-black no-underline
          min-w-max text-center"
            onClick={() =>
              window.open("https://github.com/sanjayheaven/go-gin-boilerplate")
            }
          >
            GitHub
          </div>
        </div>
      </div>

      <main>
        {/*  */}
        <div className=" my-20">
          <div className=" text-4xl font-bold text-center mb-10">
            Why use this boilerplate?
          </div>
          <div className=" grid grid-cols-3 justify-items-center w-max mx-auto gap-6">
            {features.map((item) => {
              return (
                <div
                  key={item.title}
                  className=" cursor-pointer shadow-lg rounded-lg w-[250px] p-8 transition-all duration-300"
                >
                  <div className=" font-bold ">{item.title}</div>
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
