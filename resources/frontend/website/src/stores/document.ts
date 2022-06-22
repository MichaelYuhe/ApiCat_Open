import { traverseTree } from '@natosoft/shared'
import { DOCUMENT_TYPES } from '@/common/constant'
import { defineStore } from 'pinia'
import { treeList } from '@/api/dir'

interface State {
    apiDocTree: any
}

export const extendDocTreeFeild = (node = {} as any, type = DOCUMENT_TYPES.DOC) => ({
    ...node,
    type: node.type === undefined ? type : node.type,
    isEditable: false,
    isCurrent: false,
    isLeaf: (node.type === undefined ? type : node.type) === DOCUMENT_TYPES.DOC,
})

export const useDocumentStore = defineStore({
    id: 'document',

    state: (): State => ({
        apiDocTree: [],
    }),

    getters: {},

    actions: {
        async getApiDocTree(project_id: string) {
            if (!project_id) {
                return []
            }

            try {
                const { data } = await treeList(project_id)
                this.apiDocTree = traverseTree(
                    (item: any) => {
                        item.isLeaf = item.type === DOCUMENT_TYPES.DOC
                        item.isEditable = false
                        item.isCurrent = false
                        return item
                    },
                    data || [],
                    { subKey: 'sub_nodes' }
                )
            } catch (error) {
                //
            }

            return this.apiDocTree
        },
    },
})
