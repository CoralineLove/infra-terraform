// types.ts

export interface TerraformVariable {
  type: string;
  description: string;
  default?: any;
  nullable?: boolean;
  sensitive?: boolean;
}

export interface TerraformOutput {
  value: any;
  description: string;
  sensitive?: boolean;
}

export interface TerraformResource {
  type: string;
  name: string;
  [key: string]: any; // Allows generic resource properties
}

export interface TerraformModule {
  source: string;
  version?: string; // Optional module version
  [key: string]: any; // Allows generic module properties
}

export interface TerraformProvider {
  alias?: string;
  [key: string]: any; // Allows generic provider properties
}

export interface TerraformBackend {
  type: string;
  config: {
    [key: string]: any;
  };
}

export interface TerraformConfiguration {
  terraform?: {
    required_providers?: {
      [key: string]: {
        source: string;
        version: string;
      };
    };
    required_version?: string;
  };
  provider?: TerraformProvider[];
  variable?: {
    [key: string]: TerraformVariable;
  };
  resource?: TerraformResource[];
  module?: {
    [key: string]: TerraformModule;
  };
  output?: {
    [key: string]: TerraformOutput;
  };
  terraformBackend?: TerraformBackend;
}

export type TerraformBlock = TerraformVariable | TerraformResource | TerraformModule | TerraformProvider | TerraformOutput;

export interface TerraformState {
  version: number;
  terraform_version: string;
  serial: number;
  lineage: string;
  outputs: {
    [key: string]: {
      sensitive: boolean;
      type: string;
      value: any;
    };
  };
  resources: {
    mode: string;
    type: string;
    name: string;
    provider: string;
    instances: {
      schema_version: number;
      attributes: {
        [key: string]: any;
      };
      sensitive_attributes: string[];
    }[];
  }[];
}