import {BaseService} from "./base-service";

class Config {
    subnet: string = "";
    ip_range: string = "";
    gateway: string = "";
    aux_address: string = ""
}

class IPAM {
    driver: string = "";
    config: Array<Config> = []
    // options: string = ""
}

export class NetworkEntity {
    name: string = "";
    id: string = "";
    created: string = "";
    scope: string = "";
    driver: string = "";
    internal: string = "";
    attachable: string = "";
    ingress: string = "";
    ipam: IPAM = new IPAM()
}

export class NetworkService extends BaseService {
    private static instance: NetworkService;

    public static get = () => {
        if (NetworkService.instance) {
            return NetworkService.instance
        } else {
            NetworkService.instance = new NetworkService();
            return NetworkService.instance;
        }
    };

    getNetworks(hostId: number) {
        return this.cachedPost<Array<NetworkEntity>>("/api/docker/network/list", {host_id: hostId}, 1000)
    }
}
