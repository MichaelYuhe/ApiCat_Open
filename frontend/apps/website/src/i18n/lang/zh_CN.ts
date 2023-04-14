import elementPlusLocale from 'element-plus/lib/locale/lang/zh-cn'

export default {
  name: '中文',
  app: {
    common: {
      add: '添加',
      confirm: '确定',
      save: '保存',
      export: '导出',
      import: '导入',
      cancel: '取消',
      restore: '恢复',
      emptyDataTip: '暂无数据',
      delete: '删除',
      deleteTip: '删除提示',
      confirmDelete: '确认删除{msg}吗？',
    },
    table: {
      paramName: '参数名称',
      paramType: '参数类型',
      required: '必须',
      defaultValue: '默认值',
      paramDesc: '参数说明',
    },
    form: {
      serverUrl: {
        desc: '描述',
        url: '以http://或https://开始',
      },
    },
    project: {
      list: {
        title: '项目列表',
        tabTitle: '项目',
      },
      form: {
        title: '项目名称',
        desc: '项目描述',
      },
      rules: {
        title: '请输入项目名称',
        desc: '请输入项目描述信息',
      },
      createModal: {
        title: '创建项目',
        dividerLine: '从以下方式创建',
        blackProject: '空白项目',
        importProject: '导入JSON数据文件',
        importProjectTip: '支持OpenAPI2.0、3.0',
      },
      setting: {
        title: '项目管理',
        baseInfo: '项目设置',
        serverUrl: 'URL设置',
        requestParam: '公共参数设置',
        responseParam: '公共响应设置',
        export: '导出项目',
        trash: '回收站',
      },
    },
  },
  // 编辑器
  editor: {},
  elementPlusLocale,
}
