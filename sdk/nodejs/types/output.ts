// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface GetItemSection {
    fields: outputs.GetItemSectionField[];
    id: string;
    label: string;
}

export interface GetItemSectionField {
    id: string;
    label: string;
    purpose: string;
    /**
     * (Only applies to the database category) The type of database. One of ["db2" "filemaker" "msaccess" "mssql" "mysql" "oracle" "postgresql" "sqlite" "other"]
     */
    type: string;
    value: string;
}

export interface ItemPasswordRecipe {
    /**
     * Use digits [0-9] when generating the password.
     */
    digits?: boolean;
    /**
     * The length of the password to be generated.
     */
    length?: number;
    /**
     * Use letters [a-zA-Z] when generating the password.
     */
    letters?: boolean;
    /**
     * Use symbols [!@.-_*] when generating the password.
     */
    symbols?: boolean;
}

export interface ItemSection {
    /**
     * A list of custom fields in the section.
     */
    fields?: outputs.ItemSectionField[];
    /**
     * A unique identifier for the section.
     */
    id: string;
    /**
     * The label for the section.
     */
    label: string;
}

export interface ItemSectionField {
    id: string;
    label: string;
    /**
     * Password for this item.
     */
    passwordRecipe?: outputs.ItemSectionFieldPasswordRecipe;
    purpose?: string;
    /**
     * (Only applies to the database category) The type of database. One of ["db2" "filemaker" "msaccess" "mssql" "mysql" "oracle" "postgresql" "sqlite" "other"]
     */
    type?: string;
    value: string;
}

export interface ItemSectionFieldPasswordRecipe {
    /**
     * Use digits [0-9] when generating the password.
     */
    digits?: boolean;
    /**
     * The length of the password to be generated.
     */
    length?: number;
    /**
     * Use letters [a-zA-Z] when generating the password.
     */
    letters?: boolean;
    /**
     * Use symbols [!@.-_*] when generating the password.
     */
    symbols?: boolean;
}
