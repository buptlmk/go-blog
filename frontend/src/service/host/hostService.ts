import {Response} from "../response";
import {
    CoreUsageEntity,
    DiskIO,
    GpusUsageEntity,
    HostEntity,
    MemoryUsageEntity,
    NetInterface,
    NFSStatusEntity,
    ResourcesEntity
} from "./struct";
import {BaseService} from "../base-service";


export class HostService extends BaseService {
    private static instance: HostService;

    public static get = () => {
        if (HostService.instance) {
            return HostService.instance
        } else {
            HostService.instance = new HostService();
            return HostService.instance;
        }
    };

    getNetworks() {
        return this.cachedPost<Array<string>>("/api/docker/", {}, 60000)
    }

    getHostList() {
        return this.cachedPost<Array<HostEntity>>("/api/host/list", {}, 60000)
    }

    getAvailablePorts(hostId: number): Promise<Response<Array<number>>> {
        return this.cachedPost<Array<number>>("/api/host/available_ports", {host_id: hostId,})
    }

    getResources(hostId: number) {
        return this.cachedPost<ResourcesEntity>("/api/host/resources", {host_id: hostId,}, 60000)
    }

    async getHostStatus(hostId: number, num: number = 5) {
        let res = await this.cachedPost<{ core_usage: Map<string, CoreUsageEntity>, time: string }>("/api/host/stat/average", {
            host_id: hostId,
            num: num,
        });
        if (res.data) {
            res.data.core_usage = new Map(Object.entries(res.data.core_usage))
        }
        return res
    }

    getHostMemoryStatus(hostId: number) {
        return this.cachedPost<MemoryUsageEntity>("/api/host/memory", {host_id: hostId,})
    }

    async getHostGpuStatus(hostId: number) {
        let res = await this.cachedPost<GpusUsageEntity>("/api/host/gpu", {host_id: hostId,});
        if (res.data) {
            res.data.gpus = new Map(Object.entries(res.data.gpus))
        }
        return res
    }

    getNfsServerStatus(hostId: number) {
        return this.cachedPost<string>("/api/host/nfs/server/status", {host_id: hostId}, 10000)
    }

    getNFSStatus(hostId: number) {
        return this.cachedPost<Array<NFSStatusEntity>>("/api/host/nfs/status", {host_id: hostId})
    }

    async getNetworkFlow(hostId: number, num: number = 5) {
        let res = await this.cachedPost<Map<string, NetInterface>>("/api/host/network/flow", {
            host_id: hostId,
            num: num
        });
        if (res.data) {
            res.data = new Map<string, NetInterface>(Object.entries(res.data));
        }
        return res;
    }

    async getDiskIO(hostId: number, num: number = 5) {
        let res = await this.cachedPost<Map<string, DiskIO>>("/api/host/disk/io", {
            host_id: hostId,
            num: num
        });
        if (res.data) {
            res.data = new Map<string, DiskIO>(Object.entries(res.data));
        }
        return res;
    }

    restartNfsServer(hostId: number) {
        return this.cachedPost<null>("/api/host/nfs/restart", {host_id: hostId}, 0)
    }

    remountNFS(hostId: number) {
        return this.cachedPost<Array<{
            host_id: number,
            online: boolean,
            error: string
        }>>("/api/host/nfs/mount", {host_id: hostId}, 0)
    }

    remountSelf(hostId: number){
        return this.cachedPost<Array<{
            host_id: number,
            online: boolean,
            error: string
        }>>("/api/host/nfs/bind", {host_id: hostId}, 0)
    }
}
