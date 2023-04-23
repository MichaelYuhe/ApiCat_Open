import useApi from '@/hooks/useApi'
import useCommonResponseStore from '@/store/commonResponse'
import { ProjectInfo, APICatCommonResponseCustom, APICatCommonResponse } from '@/typings'
import { createCommonResponse } from '@/views/document/components/createHttpDocument'
import { uuid } from '@apicat/shared'

export const useResponseparamList = ({ id: project_id }: Pick<ProjectInfo, 'id'>) => {
  const responseParamList: Ref<APICatCommonResponseCustom[]> = ref([])

  const commonResponseStore = useCommonResponseStore()

  const [isLoading, getResponseParamListApi] = useApi(commonResponseStore.getCommonResponseList)()

  const extendResponseParamModel = (param?: Partial<APICatCommonResponseCustom>): APICatCommonResponseCustom => {
    return {
      id: param?.id ?? uuid(),
      expand: false,
      isLoaded: false,
      isLoading: false,
      ...param,
    }
  }

  const createResponseParamModel = () => {
    const response = createCommonResponse({ description: '公共响应' })
    const extendModel = extendResponseParamModel({ code: response.code, description: response.description, isLoaded: true, expand: true })
    extendModel.detail = response
    return extendModel
  }

  const handleAddParam = () => {
    responseParamList.value.unshift(createResponseParamModel())
  }

  const handleDeleteParam = async (item: APICatCommonResponseCustom, index: number) => {
    const { detail } = item

    if (detail && detail.id) {
      item.isLoading = true
      try {
        await commonResponseStore.deleteResponseParam(project_id, detail)
      } finally {
        item.isLoading = false
      }
    }

    responseParamList.value.splice(index, 1)
  }

  onMounted(async () => {
    const data: APICatCommonResponse[] = await getResponseParamListApi(project_id)
    const list: APICatCommonResponseCustom[] = data.map((item) => extendResponseParamModel(item))
    responseParamList.value = list
  })

  return {
    isLoading,
    getResponseParamListApi,
    responseParamList,

    handleAddParam,
    handleDeleteParam,
  }
}
