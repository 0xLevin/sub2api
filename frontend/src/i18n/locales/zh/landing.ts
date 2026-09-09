export default {
  batchImageGuide: {
    title: '图片批量生成',
    description: '一次提交多条提示词，任务完成后可统一下载图片结果'
  },
  // Home Page
  home: {
    viewOnGithub: '在 GitHub 上查看',
    viewDocs: '查看文档',
    docs: '文档',
    switchToLight: '切换到浅色模式',
    switchToDark: '切换到深色模式',
    dashboard: '控制台',
    login: '登录',
    getStarted: '立即开始',
    goToDashboard: '进入控制台',
    // 新增：面向用户的价值主张
    heroSubtitle: '一个密钥，畅用多个 AI 模型',
    heroDescription: '无需管理多个订阅账号，一站式接入 Claude、GPT、Gemini 等主流 AI 服务',
    tags: {
      subscriptionToApi: '订阅转 API',
      stickySession: '会话保持',
      realtimeBilling: '按量计费'
    },
    // 用户痛点区块
    painPoints: {
      title: '你是否也遇到这些问题？',
      items: {
        expensive: {
          title: '订阅费用高',
          desc: '每个 AI 服务都要单独订阅，每月支出越来越多'
        },
        complex: {
          title: '多账号难管理',
          desc: '不同平台的账号、密钥分散各处，管理起来很麻烦'
        },
        unstable: {
          title: '服务不稳定',
          desc: '单一账号容易触发限制，影响正常使用'
        },
        noControl: {
          title: '用量无法控制',
          desc: '不知道钱花在哪了，也无法限制团队成员的使用'
        }
      }
    },
    // 解决方案区块
    solutions: {
      title: '我们帮你解决',
      subtitle: '简单三步，开始省心使用 AI'
    },
    features: {
      unifiedGateway: '一键接入',
      unifiedGatewayDesc: '获取一个 API 密钥，即可调用所有已接入的 AI 模型，无需分别申请。',
      multiAccount: '稳定可靠',
      multiAccountDesc: '智能调度多个上游账号，自动切换和负载均衡，告别频繁报错。',
      balanceQuota: '用多少付多少',
      balanceQuotaDesc: '按实际使用量计费，支持设置配额上限，团队用量一目了然。'
    },
    // 优势对比
    comparison: {
      title: '为什么选择我们？',
      headers: {
        feature: '对比项',
        official: '官方订阅',
        us: '本平台'
      },
      items: {
        pricing: {
          feature: '付费方式',
          official: '固定月费，用不完也付',
          us: '按量付费，用多少付多少'
        },
        models: {
          feature: '模型选择',
          official: '单一服务商',
          us: '多模型随意切换'
        },
        management: {
          feature: '账号管理',
          official: '每个服务单独管理',
          us: '统一密钥，一站管理'
        },
        stability: {
          feature: '服务稳定性',
          official: '单账号易触发限制',
          us: '多账号池，自动切换'
        },
        control: {
          feature: '用量控制',
          official: '无法限制',
          us: '可设配额、查明细'
        }
      }
    },
    providers: {
      title: '已支持的 AI 模型',
      description: '一个 API，多种选择',
      supported: '已支持',
      soon: '即将推出',
      claude: 'Claude',
      gemini: 'Gemini',
      antigravity: 'API',
      more: '更多'
    },
    // CTA 区块
    cta: {
      title: '准备好开始了吗？',
      description: '注册即可获得免费试用额度，体验一站式 AI 服务',
      button: '免费注册'
    },
    v2: {
      kicker: 'OPENAI API 接入',
      buying: {
        kicker: '购买与接入',
        title: '像买云资源一样购买 OpenAI Token',
        description: '充值后即可生成 API Key，余额、消耗和调用记录集中可查，适合快速接入 GPT 能力。',
        items: {
          topUp: {
            label: 'Step 01',
            title: '充值额度',
            description: '选择合适的充值金额，余额到账后即可开始调用，不需要先理解复杂套餐。'
          },
          key: {
            label: 'Step 02',
            title: '获取 API Key',
            description: '为个人项目、团队成员或不同业务分别生成 Key，泄露或停用时也容易处理。'
          },
          usage: {
            label: 'Step 03',
            title: '按量消费',
            description: '每次请求都会扣减余额并保留记录，客户可以随时查看 GPT 调用、Token 和费用明细。'
          }
        }
      },
      capabilities: {
        kicker: '你能获得什么',
        title: 'OpenAI Token 购买与调用',
        description: '从购买额度到发起请求保持简单：充值、拿 Key、接入应用，并随时查看余额变化。',
        items: {
          api: {
            title: '兼容 API',
            description: '用熟悉的 OpenAI 兼容请求方式接入，迁移脚本、插件和产品后端更轻松。'
          },
          routing: {
            title: 'GPT 场景覆盖',
            description: '围绕文本生成、代码辅助和自动化调用场景使用 OpenAI Token。'
          },
          accountPool: {
            title: '充值即用',
            description: '购买 Token 后即可使用，无需预付多份订阅，也不用维护复杂账号。'
          },
          billing: {
            title: '实时扣费',
            description: '每次请求都会记录消耗，余额和消费明细清楚展示。'
          },
          guardrails: {
            title: '额度控制',
            description: '可以为 API Key 设置使用限制，避免脚本或团队成员意外超支。'
          },
          observability: {
            title: '用量查询',
            description: '按时间、GPT 调用和 Key 查看请求记录，方便对账和排查问题。'
          }
        }
      },
      workflow: {
        kicker: '使用流程',
        title: '四步开始消耗 Token',
        description: '从注册到第一次调用，流程保持简单；你把注意力放在业务和提示词上就好。',
        steps: {
          connect: {
            title: '注册账号',
            description: '创建账户并进入控制台，查看余额、密钥和用量入口。'
          },
          issue: {
            title: '充值 Token',
            description: '按需购买 OpenAI Token 额度，不必为了临时调用承担固定订阅。'
          },
          route: {
            title: '创建 API Key',
            description: '为项目、脚本或团队成员分别创建 Key，方便管理和停用。'
          },
          measure: {
            title: '开始调用',
            description: '把 Key 放进你的应用，按实际请求扣除 Token，并随时查看明细。'
          }
        }
      },
      controlPlane: {
        kicker: '账户视图',
        metrics: {
          access: '接入',
          models: '服务商',
          usage: '用量',
          balance: '余额'
        },
        signals: {
          failover: '一个 API Key 即可接入 OpenAI 兼容调用',
          usage: '请求、Token 和费用明细实时记录',
          audit: '按项目或团队成员拆分 Key，管理更清楚'
        }
      },
      modelAccess: {
        kicker: 'OpenAI 能力',
        title: '一个余额覆盖常见 GPT 调用场景',
        description: '统一处理 OpenAI Token 购买、API Key 和用量记录，让日常 GPT 调用更容易管理。',
        items: {
          chat: {
            title: '聊天与内容生成',
            description: '用于客服回复、内容草稿、资料整理、知识库问答等高频文本任务。'
          },
          code: {
            title: '代码与开发工具',
            description: '用于脚本生成、代码解释、测试补全、IDE 插件或内部开发助手。'
          },
          automation: {
            title: '自动化工作流',
            description: '用于定时任务、机器人、数据处理链路和产品后端的 GPT 调用。'
          }
        }
      },
      useCases: {
        kicker: '适用场景',
        title: '适合需要稳定消耗 OpenAI Token 的用户',
        items: {
          teams: {
            title: '团队协作',
            description: '给不同成员或项目分配 Key，查看各自的 Token 消耗。'
          },
          products: {
            title: '产品接入',
            description: '把 GPT 能力接入 SaaS、机器人、插件或自动化工作流。'
          },
          resellers: {
            title: '个人与开发者',
            description: '不用维护多份订阅，用统一余额完成实验、开发和日常调用。'
          }
        }
      },
      faq: {
        kicker: '常见问题',
        title: '充值前需要确认的几件事',
        items: {
          billing: {
            question: 'Token 怎么扣费？',
            answer: '按实际请求消耗扣减余额，具体消耗会进入用量记录，方便查询和对账。'
          },
          compatibility: {
            question: '现有应用能接入吗？',
            answer: '通过兼容 API 和 API Key 接入，适合脚本、插件、后端服务和自动化工具。'
          },
          balance: {
            question: '余额能看到明细吗？',
            answer: '可以查看请求记录、Token 消耗、费用和余额变化，避免黑盒消费。'
          },
          keys: {
            question: '能给不同项目分 Key 吗？',
            answer: '可以为不同项目、成员或环境创建独立 Key，便于限制、停用和追踪使用情况。'
          }
        }
      },
      finalCta: {
        kicker: '开始使用'
      }
    },
    footer: {
      allRightsReserved: '保留所有权利。'
    }
  },

  // Key Usage Query Page
  keyUsage: {
    title: 'API Key 用量查询',
    subtitle: '输入您的 API Key 以查看实时消费金额与使用状态',
    placeholder: 'sk-ant-mirror-xxxxxxxxxxxx',
    query: '查询',
    querying: '查询中...',
    privacyNote: '您的 Key 仅在浏览器本地处理，不会被存储',
    dateRange: '统计范围:',
    dateRangeToday: '今日',
    dateRange7d: '7 天',
    dateRange30d: '30 天',
    dateRange90d: '90 天',
    dateRangeCustom: '自定义',
    apply: '应用',
    used: '已使用',
    detailInfo: '详细信息',
    tokenStats: 'Token 统计',
    dailyDetail: '按日明细',
    modelStats: '模型用量统计',
    // Table headers
    date: '日期',
    model: '模型',
    requests: '请求数',
    inputTokens: '输入 Tokens',
    outputTokens: '输出 Tokens',
    cacheCreationTokens: '缓存创建',
    cacheReadTokens: '缓存读取',
    cacheWriteTokens: '缓存写入',
    totalTokens: '总 Tokens',
    cost: '费用',
    // Status
    quotaMode: 'Key 限额模式',
    walletBalance: '钱包余额',
    // Ring card titles
    totalQuota: '总额度',
    limit5h: '5 小时限额',
    limitDaily: '日限额',
    limit7d: '7 天限额',
    limitWeekly: '周限额',
    limitMonthly: '月限额',
    // Detail rows
    remainingQuota: '剩余额度',
    expiresAt: '过期时间',
    todayExpires: '(今日到期)',
    daysLeft: '({days} 天)',
    usedQuota: '已用额度',
    resetNow: '即将重置',
    subscriptionType: '订阅类型',
    subscriptionExpires: '订阅到期',
    // Usage stat cells
    todayRequests: '今日请求',
    todayInputTokens: '今日输入',
    todayOutputTokens: '今日输出',
    todayTokens: '今日 Tokens',
    todayCacheCreation: '今日缓存创建',
    todayCacheRead: '今日缓存读取',
    todayCost: '今日费用',
    rpmTpm: 'RPM / TPM',
    totalRequests: '累计请求',
    totalInputTokens: '累计输入',
    totalOutputTokens: '累计输出',
    totalTokensLabel: '累计 Tokens',
    totalCacheCreation: '累计缓存创建',
    totalCacheRead: '累计缓存读取',
    totalCost: '累计费用',
    avgDuration: '平均耗时',
    // Messages
    enterApiKey: '请输入 API Key',
    querySuccess: '查询成功',
    queryFailed: '查询失败',
    queryFailedRetry: '查询失败，请稍后重试',
    noDailyUsage: '暂无按日用量数据',
  },

  // Setup Wizard
  setup: {
    title: 'Sub2API 安装向导',
    description: '配置您的 Sub2API 实例',
    database: {
      title: '数据库配置',
      description: '连接到您的 PostgreSQL 数据库',
      host: '主机',
      port: '端口',
      username: '用户名',
      password: '密码',
      databaseName: '数据库名称',
      sslMode: 'SSL 模式',
      passwordPlaceholder: '密码',
      ssl: {
        disable: '禁用',
        require: '要求',
        verifyCa: '验证 CA',
        verifyFull: '完全验证'
      }
    },
    redis: {
      title: 'Redis 配置',
      description: '连接到您的 Redis 服务器',
      host: '主机',
      port: '端口',
      username: '用户名（可选）',
      password: '密码（可选）',
      database: '数据库',
      usernamePlaceholder: '默认用户留空',
      passwordPlaceholder: '密码',
      enableTls: '启用 TLS',
      enableTlsHint: '连接 Redis 时使用 TLS（公共 CA 证书）'
    },
    admin: {
      title: '管理员账户',
      description: '创建您的管理员账户',
      email: '邮箱',
      password: '密码',
      confirmPassword: '确认密码',
      passwordPlaceholder: '至少 8 个字符',
      confirmPasswordPlaceholder: '确认密码',
      passwordMismatch: '密码不匹配'
    },
    ready: {
      title: '准备安装',
      description: '检查您的配置并完成安装',
      database: '数据库',
      redis: 'Redis',
      adminEmail: '管理员邮箱'
    },
    status: {
      testing: '测试中...',
      success: '连接成功',
      testConnection: '测试连接',
      installing: '安装中...',
      completeInstallation: '完成安装',
      completed: '安装完成！',
      redirecting: '正在跳转到登录页面...',
      restarting: '服务正在重启，请稍候...',
      timeout: '服务重启时间超出预期，请手动刷新页面。'
    }
  },

  // Common
}
